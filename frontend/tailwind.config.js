/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        bg: '#FAF7F2',
        primary: '#C8B6A6',
        secondary: '#B0BBA2',
        text: '#3E362E',
        card: '#FFFFFF',
      }
    },
  },
  plugins: [],
}