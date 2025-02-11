/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./render/views/**/*.templ"],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        open: ['"Open Sans"', "sans-serif"],
      },
      colors: {
        pinky: "#2b222a",
        purpur: "#d164a9",
        bluey: "#4477CE",
        lightBluey: "#8CABFF",
      },
      screens: {
        ph: "360px",
      },
    },
  },
  plugins: [],
};
