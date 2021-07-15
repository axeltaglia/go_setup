export function resolveErrorMessage(e) {
    if (e.response && e.response.data && e.response.data.errorMessage) {
        return e.response.data.errorMessage
    } else {
        let errorString = e.toString();
        if (errorString === "") {
            errorString = "" + e;
        }
        return "An unexpected error has occurred (" + errorString + ")";
    }
}