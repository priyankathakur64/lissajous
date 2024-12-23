# Lissajous Multi-Color Animation

This project is a **Golang-based HTTP server** that generates animated **Lissajous figures** in GIF format. The figures are multi-colored, showcasing beautiful oscillatory patterns based on mathematical equations.

## **Features**
- Dynamically generates multi-colored Lissajous animations.
- Handles HTTP requests to display the animation in a browser.
- Each frame in the GIF cycles through different colors in the palette.


## **How It Works**
1. The server listens on port `8080`.
2. When accessed via the browser (e.g., `http://localhost:8080`), it generates and displays a multi-colored Lissajous animation as a GIF.
3. The animation is generated using:
   - A set of sinusoidal equations.
   - Randomly assigned colors for each frame.

## **Prerequisites**
- Go programming language installed (version 1.16+ recommended).




