/**
 * HTMX utils
 */
document.addEventListener("htmx:responseError", (event) => {
  const xhr = event.detail.xhr;

  if (xhr.status === 400) {
    const errorData = JSON.parse(xhr.responseText);

    const errorContainer = document.querySelector("#error");
    const errorElement = document.createElement("p");
    errorElement.textContent = errorData.error;
    errorContainer.textContent = "";
    errorContainer.appendChild(errorElement);
  }
});
// End HTMX utils

/**
 * App Utils
 */
function copyToClipboard() {
  const text = document.querySelector("#shortened-url").value;
  console.log(text);
  navigator.clipboard.writeText(text);
}

function resetForm() {
  document.querySelector("#error").textContent = "";
  document.querySelector("#shortened-url").value = "";
}
// End App Utils

/**
 * App utils
 */
function copyToClipboard() {
  const text = document.querySelector("#shortened-url").value;
  console.log(text);
  navigator.clipboard.writeText(text);
}

function resetForm() {
  document.querySelector("#error").textContent = "";
  document.querySelector("#shortened-url").value = "";
}
// End app utils
