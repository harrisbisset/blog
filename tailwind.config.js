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
        screens: {
          'ph': '360px',
          // => @media (min-width: 360px) { ... }
    
          'sm': '640px',
          // => @media (min-width: 640px) { ... }
    
          'xl': '1240px',
          // => @media (min-width: 1280px) { ... }
    
        },
      },
    },
    plugins: [],
    corePlugins: {
      preflight: true,
    }
  }