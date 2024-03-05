/**
* Piece dragging funcitonality.
* @param {DragEvent} e
*/
function dragPiece(e) {
    e.dataTransfer.setData("text", e.target.id);
}

/**
* Allow pieces to be dropped on squares
* @param {DragEvent} e
*/
function allowDrop(e) {
    e.preventDefault();
}

/**
* Transfer piece to new square
* @param {DragEvent} e
*/
function drop(e) {
    e.preventDefault();
    const srcId = e.dataTransfer.getData("text");
    const srcPiece = document.getElementById(srcId);

    console.log(e.target.className);
    if (typeof e.target.className === "string" &&
        e.target.className.indexOf("square") !== -1) {
        e.target.appendChild(srcPiece);
    } else {
        destPiece = e.target;
        while (!destPiece.id) {
            destPiece = destPiece.parentNode;
        }
        if (destPiece.id === srcId) { return; }

        square = destPiece.parentNode;
        while (square.firstChild) {
            square.removeChild(square.lastChild);
        }
        square.appendChild(srcPiece);
    }
}
