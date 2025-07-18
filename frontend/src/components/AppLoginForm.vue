<template>
  <Form :validation-schema="schema" @submit="onSubmit">
    <div class="flex flex-col gap-6">
      <ErrorAlert :message="errorMsg" @clear="errorMsg = ''" />
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
      <button
        type="submit"
        class="text-[#fff] text-lg font-bold py-2 px-8 rounded-xl bg-[#055a32]"
        :disabled="loading"
      >
        {{ loading ? "Loggin in..." : "Login to your account" }}
      </button>
    </div>
  </Form>
</template>

<script setup>
import { ref } from "vue";
import { Form, Field, ErrorMessage, useForm } from "vee-validate";
import * as yup from "yup";
import api from "@/api";
import ErrorAlert from "./ErrorAlert.vue";

import { useModalStore } from "@/stores/modal";
import { useUserStore } from "@/stores/user";

const modalStore = useModalStore();
const userStore = useUserStore();

const errorMsg = ref("");
const loading = ref(false);

const schema = yup.object({
  email: yup
    .string()
    .email("Must be a valid email")
    .required("Email is required"),
  password: yup
    .string()
    .min(6, "Password must be at least 6 characters")
    .required("Password is required"),
});

const onSubmit = async (values) => {
  loading.value = true;
  try {
    await userStore.login(values.email, values.password);
    modalStore.close();
  } catch (err) {
    errorMsg.value = err.response?.data?.message || "Invalid login credentials";
    email.value = "";
    password.value = "";
    console.error("API error", err);
  } finally {
    loading.value = false;
  }
};
</script>
