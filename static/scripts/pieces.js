/**
* Piece dragging funcitonality.
* @param {object} e
*/
function dragPiece(e) {
    e.dataTransfer.setData("text", e.target.id);
}

/**
* Allow pieces to be dropped on squares
* @param {object} e
*/
function allowDrop(e) {
    e.preventDefault();
}

/**
* Transfer piece to new square
* @param {object} e
*/
function drop(e) {
    e.preventDefault();
    const data = e.dataTransfer.getData("text");
    const el = document.getElementById(data);
    e.target.appendChild(el);
}
