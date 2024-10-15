/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./frontend/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ["emerald"]
  }
}

