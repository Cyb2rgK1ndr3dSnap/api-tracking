document.addEventListener('DOMContentLoaded', function() {
    const mobileToggle = document.getElementById('mobileToggle');
    const desktopToggle = document.getElementById('desktopToggle');
    const sidebar = document.getElementById('sidebar');
    const content = document.getElementById('content');
    const spinner = document.getElementById('spinner');
    const toggleIcon = document.querySelector('.toggle-icon');
    const navItems = document.querySelectorAll('.nav-item');

    function toggleSidebar() {
        sidebar.classList.toggle('active');
        sidebar.classList.toggle('collapsed');
        content.classList.toggle('sidebar-active');
        desktopToggle.textContent = sidebar.classList.contains('active') ? '◀' : '▶';
    }

    // Handle mobile toggle click
    mobileToggle.addEventListener('click', function() {
        if (window.innerWidth <= 768) {
            // Show loading spinner
            toggleIcon.style.display = 'none';

                toggleIcon.style.display = 'block';
                toggleSidebar();
        }
    });

    // Handle desktop toggle click
    desktopToggle.addEventListener('click', function() {
        if (window.innerWidth > 768) {
            toggleSidebar();
        }
    });

    // Handle submenu toggles
    navItems.forEach(item => {
        const link = item.querySelector('.nav-link');
        if (item.querySelector('.submenu')) {
            link.addEventListener('click', function(e) {
                e.preventDefault();
                
                // Close other open submenus
                navItems.forEach(otherItem => {
                    if (otherItem !== item && otherItem.classList.contains('active')) {
                        otherItem.classList.remove('active');
                    }
                });

                // Toggle current submenu
                item.classList.toggle('active');
            });
        }
    });

    // Close sidebar when clicking outside on mobile
    document.addEventListener('click', function(e) {
        if (window.innerWidth <= 768) {
            if (!sidebar.contains(e.target) && !mobileToggle.contains(e.target)) {
                sidebar.classList.remove('active');
            }
        }
    });
});