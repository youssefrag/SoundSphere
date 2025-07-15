<template>
  <Form :validation-schema="schema" @submit="onSubmit">
    <div class="flex flex-col gap-6">
      <div>
        <label for="name" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
          >Name</label
        >
        <Field
          id="name"
          name="name"
          as="input"
          type="text"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="John Smith"
        />
        <ErrorMessage name="name" class="text-red-600 text-sm mt-1" />
      </div>
      <div>
        <label for="email" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
          >Email</label
        >
        <Field
          id="email"
          name="email"
          as="input"
          type="email"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="john@example.com"
        />
        <ErrorMessage name="email" class="text-red-600 text-sm mt-1" />
      </div>

      <div>
        <label
          for="password"
          class="block mb-2 text-lg font-bold text-[#0E0E0F]"
          >Password</label
        >
        <Field
          id="password"
          name="password"
          as="input"
          type="password"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="password"
        />
        <ErrorMessage name="password" class="text-red-600 text-sm mt-1" />
      </div>

      <div>
        <label
          for="confirmPassword"
          class="block mb-2 text-lg font-bold text-[#0E0E0F]"
          >Confirm Password</label
        >
        <Field
          id="confirmPassword"
          name="confirmPassword"
          as="input"
          type="password"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="confirm password"
        />
        <ErrorMessage
          name="confirmPassword"
          class="text-red-600 text-sm mt-1"
        />
      </div>

      <input
        ref="fileInput"
        type="file"
        accept="image/*"
        @change="onFileChange"
        class="hidden"
      />

      <button
        type="button"
        class="text-white text-lg font-bold py-2 px-8 rounded-xl bg-gray-700"
        @click="triggerFileDialog"
      >
        <font-awesome-icon :icon="['fas', 'upload']" class="text-[#FFA900]" />
        {{ fileName || "Upload Picture" }}
      </button>
      <button
        type="submit"
        class="text-[#fff] text-lg font-bold py-2 px-8 rounded-xl bg-[#055a32]"
        :disabled="loading"
      >
        {{ loading ? "Registering…" : "Register your account" }}
      </button>
    </div>
  </Form>
</template>
<script setup>
import { ref as vueRef } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

import { useModalStore } from "@/stores/modal";
import { useUserStore } from "@/stores/user";

import { storage, storageRef, uploadBytesToPath, getURL } from "@/firebase";

const modalStore = useModalStore();
const userStore = useUserStore();

const file = vueRef(null);
const fileName = vueRef("");
const loading = vueRef(false);
const fileInput = vueRef(null);

function triggerFileDialog() {
  fileInput.value?.click();
}

function onFileChange(e) {
  const selected = e.target.files?.[0];
  if (selected) {
    file.value = selected;
    fileName.value = selected.name;
  }
}

async function uploadAvatar(file) {
  const path = `avatars/${file.name}`;
  const avatarRef = storageRef(storage, path);
  await uploadBytesToPath(avatarRef, file);
  return await getURL(avatarRef);
}

const schema = yup.object({
  name: yup.string().required("Name is required"),
  email: yup
    .string()
    .email("Must be a valid email")
    .required("Email is required"),
  password: yup
    .string()
    .min(6, "Password must be at least 6 characters")
    .required("Password is required"),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref("password")], "Passwords must match")
    .required("Please confirm your password"),
});

async function onSubmit(values) {
  loading.value = true;
  try {
    const avatarUrl = file.value ? await uploadAvatar(file.value) : undefined;

    await userStore.register(
      values.name,
      values.email,
      values.password,
      avatarUrl
    );
    modalStore.close();
  } catch (err) {
    console.error(err);
    alert("Registration failed: " + (err.response?.data?.error || err.message));
  } finally {
    loading.value = false;
  }
}
</script>
