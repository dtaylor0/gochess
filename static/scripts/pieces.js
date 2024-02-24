document.querySelectorAll('.chess-piece').forEach(function(item) {
    item.addEventListener('dragstart', function(event) {
        console.log('drag started');
    });
});


/**
* Piece dragging funcitonality.
* @param {object} e
*/
function dragPiece(e) {
    console.log("drag started")
}
