/** @type {import('tailwindcss').Config} */
export const content = ['./index.html', './src/**/*.{js,ts,jsx,tsx}'];
export const theme = {
  extend: {
    colors: {
      // Define your custom color here
      buttonBlue: '#007bff',
    },
  },
};
// export const plugins = [
//   require('@tailwindcss/forms'),
//   // require('@tailwindcss/aspect-ratio'),
// ];
