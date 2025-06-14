document.addEventListener("DOMContentLoaded", function () {
  const textarea = document.querySelector("textarea");

  const savedWidth = localStorage.getItem("textareaWidth");
  const savedHeight = localStorage.getItem("textareaHeight");
  if (savedWidth) textarea.style.width = savedWidth;
  if (savedHeight) textarea.style.height = savedHeight;

  textarea.addEventListener("mouseup", function () {
    localStorage.setItem("textareaWidth", textarea.style.width);
    localStorage.setItem("textareaHeight", textarea.style.height);
  });
});
