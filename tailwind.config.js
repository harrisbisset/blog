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
          bluey: "#4477CE",
          lightBluey: "#8CABFF",
        },
      },
    },
    plugins: [],
    corePlugins: {
      preflight: true,
    }
  }