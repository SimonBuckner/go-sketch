package main

import "math"

func ProjectIsoXY(originX, originY, tileX, tileY, tileW, tileH float64) (x, y float64) {

	w := tileW * 0.5
	h := tileH * 0.5

	x = math.Floor(originX + (tileX * w) - (tileY * w))
	y = math.Floor(originY + (tileX * h) + (tileY * h))

	return
}

//   function renderFilledTile(w, h, strokeColor, fillColor, ...vertices) {
// 	var tCanvas = document.createElement('canvas');
// 	var tContext = tCanvas.getContext('2d');
// 	tContext.width = w;
// 	tContext.height = h;

// 	tContext.translate(w * 0.5, h * 0.5);
// 	fillPoly(tContext, strokeColor, fillColor, w, h, ...vertices);
// 	return { canvas: tCanvas, context: tContext };
//   };

//   function renderEmptyTile(w, h, strokeColor, ...vertices) {
// 	var tCanvas = document.createElement('canvas');
// 	var tContext = tCanvas.getContext('2d');
// 	tContext.width = w;
// 	tContext.height = h;

// 	tContext.translate(w * 0.5, h * 0.5);
// 	drawPoly(tContext, strokeColor,   w, h, ...vertices);
// 	return { canvas: tCanvas, context: tContext };
//   };

//   function drawPoint(context, color, x, y) {
// 	context.strokeStyle = color;
// 	context.beginPath();
// 	context.rect(x, y, 1, 1);
// 	context.stroke();
//   };

//   function drawPoly(context, color, w, h, ...vertices) {
// 	context.strokeStyle = color;
// 	context.beginPath();

// 	if (vertices.length < 2) {
// 	  context.rect(x, y, 1, 1);
// 	  context.stroke();
// 	  return;
// 	};

// 	const x = vertices[0][0] * w * 0.5;
// 	const y = vertices[0][1] * h * 0.5;

// 	context.beginPath();
// 	context.moveTo(x, y);
// 	for (let i = 1; i < vertices.length; i++) {
// 	  const x = vertices[i][0] * w * 0.5;
// 	  const y = vertices[i][1] * h * 0.5;
// 	  context.lineTo(x, y);
// 	}
// 	context.stroke();
//   };

//   function fillPoly(context, strokeColor, fillColor, w, h, ...vertices) {

// 	if (vertices.length < 2) { return; };

// 	const x = vertices[0][0] * w * 0.5;
// 	const y = vertices[0][1] * h * 0.5;

// 	context.beginPath();
// 	context.fillStyle = fillColor;

// 	context.moveTo(x, y);
// 	for (let i = 1; i < vertices.length; i++) {
// 	  const x = vertices[i][0] * w * 0.5;
// 	  const y = vertices[i][1] * h * 0.5;
// 	  context.lineTo(x, y);
// 	}
// 	context.fill();
// 	context.linewidth = 1;
// 	context.strokeStyle = strokeColor;
// 	context.stroke();

//   };

//   function rotate2D(angle, ...vertices) {
// 	let rotated = [];
// 	const rad = angle * Math.PI / 180;

// 	for (let i = 0; i < vertices.length; i++) {
// 	  const x = vertices[i][0] * Math.cos(rad) - vertices[i][1] * Math.sin(rad);
// 	  const y = vertices[i][1] * Math.cos(rad) + vertices[i][0] * Math.sin(rad);
// 	  rotated.push([x, y]);
// 	};
// 	return rotated;
//   };

//   export {
// 	isoProjectXY,
// 	renderFilledTile,
// 	renderEmptyTile,
// 	drawPoly,
// 	fillPoly,
// 	rotate2D,
// 	drawPoint,
//   };
