/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      'view/**/**/*.templ',
    ],
    darkMode: 'class',
    theme: {
      extend: {
        fontFamily: {
          open: ['"Open Sans"', 'sans-serif'],
        },
        colors: {
          pinky: "#2b222a",
          purpur: "#d164a9",
        },
      },
    },
    plugins: [],
    corePlugins: {
      preflight: true,
    }
  }