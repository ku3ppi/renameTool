// Select elements
const draggable = document.getElementById("draggable-item");
const dropContainer = document.getElementById("drag-container");

// Add event listeners for the draggable item
draggable.addEventListener("dragstart", (e) => {
    e.dataTransfer.setData("text/plain", draggable.id); // Pass the item's ID
    draggable.style.opacity = "0.5"; // Visual feedback
});

draggable.addEventListener("dragend", () => {
    draggable.style.opacity = "1"; // Reset opacity
});

// Add event listeners for the drop container
dropContainer.addEventListener("dragover", (e) => {
    e.preventDefault(); // Necessary to allow dropping
    dropContainer.style.borderColor = "#007bff"; // Highlight border
});

dropContainer.addEventListener("dragleave", () => {
    dropContainer.style.borderColor = "#aaa"; // Reset border color
});

dropContainer.addEventListener("drop", (e) => {
    e.preventDefault(); // Prevent default drop behavior
    dropContainer.style.borderColor = "#aaa"; // Reset border color

    // Get the ID of the dragged element and append it to the drop container
    const draggableId = e.dataTransfer.getData("text/plain");
    const draggedElement = document.getElementById(draggableId);
    dropContainer.appendChild(draggedElement);
});
