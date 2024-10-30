function setColorScheme(scheme) {
    document.documentElement.classList.add(scheme);
    return
}
  
function getPreferredColorScheme() {
    if (window.matchMedia) {
        if(window.matchMedia('(prefers-color-scheme: dark)').matches){
            return 'dark';
        } else {
            return 'light';
        }
    }
    return 'light';
}
  
function updateColorScheme(){
    setColorScheme(getPreferredColorScheme());
    return
}
  
if(window.matchMedia){
    var colorSchemeQuery = window.matchMedia('(prefers-color-scheme: dark)');
    colorSchemeQuery.addEventListener('change', updateColorScheme);
}

updateColorScheme();