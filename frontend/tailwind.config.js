/** @type {import('tailwindcss').Config} */
export default {
  purge: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  content: [],
  theme: {
    extend: {
      screens: {
        // everything ≤1023px → “medium-and-smaller”
        "md-down": { max: "1023px" },
        // everything ≤767px  → “small-and-smaller”
        "sm-down": { max: "767px" },
      },
    },
  },
  plugins: [],
};
