// validation-schema.js
import { object, string, ref as yupRef } from "yup";

export const registrationSchema = object({
  name: string()
    .required("Name is required")
    .min(2, "Name must be at least 2 characters"),
  email: string().required("Email is required").email("Must be a valid email"),
  password: string()
    .required("Password is required")
    .min(6, "Password must be at least 6 characters"),
  confirmPassword: string()
    .required("Please confirm your password")
    .oneOf([yupRef("password")], "Passwords do not match"),
});
