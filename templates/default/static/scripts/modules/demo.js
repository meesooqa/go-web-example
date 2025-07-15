function showAlert() {
    alert("JavaScript работает!");
    document.getElementById("demo").textContent = "Текст изменен через JavaScript!";
}

export function initDemo() {
    const btn = document.getElementById('btnDemo');
    if (btn) {
        btn.addEventListener('click', showAlert);
    }
}
