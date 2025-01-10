const loginForm = document.getElementById('login-form');

    loginForm.addEventListener('submit', async (event) => {
        event.preventDefault(); // Prevenir recarga de la página

        const formData = new FormData(loginForm);
        const data = {
            username: formData.get('username'),
            password: formData.get('password')
        };

        try {
            const response = await fetch('/user/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                const result = await response.json();
                window.location.href = '/index';
            } else {
                const error = await response.json();
                alert(`Login failed: ${error.error}`);
            }
        } catch (err) {
            console.error('Error:', err);
            alert('An error occurred during login.');
        }
    });