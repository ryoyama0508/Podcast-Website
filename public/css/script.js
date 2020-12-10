const hamburger = document.getElementById('hamburger');
const nav = document.getElementById('nav-ul');

hamburger.addEventListener('click', () => {
    console.log('here');
    nav.classList.toggle('show');
});