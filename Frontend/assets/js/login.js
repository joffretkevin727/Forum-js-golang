const boxSignIn = document.querySelector('.box.active');
const boxSignUp = document.querySelector('.sign-up');

boxSignUp.style.display = 'none';

document.querySelectorAll('.row-start').forEach(rowStart => {
    rowStart.querySelectorAll('p').forEach((tab, index) => {
        tab.addEventListener('click', () => {
            document.querySelectorAll('.row-start p').forEach(p => p.classList.remove('active'));
            document.querySelectorAll(`.row-start p:nth-child(${index + 1})`).forEach(p => p.classList.add('active'));

            if (index === 0) {
                boxSignIn.style.display = 'block';
                boxSignUp.style.display = 'none';
            } else {
                boxSignIn.style.display = 'none';
                boxSignUp.style.display = 'block';
            }
        });
    });
});