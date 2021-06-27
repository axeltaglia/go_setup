export function loadAuthToken() {
    return window.localStorage.getItem("pepToken");
}

export function persistAuthToken(value) {
    if (value) {
        window.localStorage.setItem("pepToken", value);
    } else {
        window.localStorage.removeItem("pepToken");
    }
}