/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./templates/**/*'],
    theme: {
        fontFamily: {
            'open-sans': ['"Open Sans"', 'sans-serif'],
            montserrat: ['Montserrat', 'sans-serif'],
            'playfair-display': ['"Playfair Display"', 'serif']
        },
        extend: {
            colors: {
                'shadow-green': {
                    50: '#f3f8f7',
                    100: '#e0edec',
                    200: '#c4dddc',
                    300: '#9bc5c3',
                    400: '#6ba5a2',
                    500: '#4f8b8a',
                    600: '#457475',
                    700: '#3c6062',
                    800: '#375153',
                    900: '#314648',
                    950: '#1d2d2f'
                },
                'chathams-blue': {
                    50: '#f1f9fe',
                    100: '#e3f1fb',
                    200: '#c0e3f7',
                    300: '#89cdf0',
                    400: '#4ab3e6',
                    500: '#239ad4',
                    600: '#157bb4',
                    700: '#126292',
                    800: '#13547a',
                    900: '#154665',
                    950: '#0e2d43'
                }
            }
        }
    },
    plugins: []
};
