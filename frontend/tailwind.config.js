/** @type {import('tailwindcss').Config} */
export default {
  purge: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  content: [],
  theme: {
    extend: {
      screens: {
        "md-down": { max: "1023px" },
        "sm-down": { max: "767px" },
      },
    },
  },
  plugins: [],
};
