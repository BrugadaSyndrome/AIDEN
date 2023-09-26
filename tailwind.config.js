/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./**/*.{html,js}"
  ],
  theme: {
    extend: {
      fontFamily: {
        "pixel-splitter": ["pixel-splitter", "ui-sans-serif"]
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

