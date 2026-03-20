/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    container: {
      center: true,
      padding: '1rem',
    },
    extend: {
      maxWidth: {
        '7xl': '1200px',
      },
      colors: {
        primary: {
          DEFAULT: '#ff4400',
          hover: '#ff5722',
          light: '#fff3e0',
        }
      }
    },
  },
  plugins: [],
}
