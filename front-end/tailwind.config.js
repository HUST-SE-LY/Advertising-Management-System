/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      keyframes: {
        shine: {
          '0%': {
            opacity: 0
          },
          '50%': {
            opacity: 1
          },
          '100%': {
            opacity: 0
          }
        },
        floatin: {
          '0%': {
            transform: 'translateY(50px)',
            opacity: 0.5,
          },
          '100%': {
            transform: 'translateY(0)',
            opacity: 1,
          }
        },
        float: {
          '0%': {
            transform: 'translateY(50px)',
          },
          '50%': {
            transform: 'translateY(-50px)'
          },
          '100%': {
            transform: 'translateY(50px)'
          }
        },
        listItemIn: {
          '0%': {
            transform: 'translateY(200px)',
            opacity: 0,
            scale: '0.5',
          },
          '100%': {
            transform: 'translateY(0px)',
            opacity: 1,
            scale: '1',
          }
        },
        fadein: {
          '0%': {
            opacity: 0,
          },
          '100%': {
            opacity: 1,
          }
        }
      },
      animation: {
        shine: 'shine 2s infinite linear',
        floatin: 'floatin 0.5s forwards',
        float: 'float 10s infinite',
        listItemIn: 'listItemIn 0.3s ease-out forwards',
        fadein: 'fadein 0.3s ease-out forwards'
      },
      boxShadow: {
        'home': '0px 0 500px rgba(0,0,0,0.1)'
      },
      screens: {
        'lg-height': {
          'raw': '(min-height: 1080px)'
        },
        'md-height': {
          'raw': '(min-height: 800px)'
        }
      }
    },
  },
  plugins: [],
};
