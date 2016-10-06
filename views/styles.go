package views

const styles = `
html, body {
    margin: 0;
    padding: 0;
}

body {
    font: 100%/1.3 Verdana, sans-serif;
    overflow-y: scroll;
}

.container {
    max-width: 40rem;
    margin: 1rem;
}

aside {
    max-width: 40rem;
    margin: 1rem 1.5rem;
}

.container:before, .container:after {
    clear: both;
    content: " ";
    display: table;
}

a {
    color: hsl(220, 51%, 44%);
}

a:hover {
    color: hsl(208, 56%, 38%);
}

header {
    margin: 0 1rem;
    font-size: 1rem;
    border-bottom: 1px solid #ddd;
}

header h1, header > a {
    margin: 0;
    padding: 1.3rem;
    height: 1.3rem;
    line-height: 1.3rem;
    display: inline-block;
}

header h1 {
    font-size: 1.5rem;
    padding-left: 0;
    margin-left: .5rem;
    font-weight: bold;
    align-self: flex-start;
}

header h1 a {
    color: #000;
    text-decoration: none;
}

.feeds {
    width: auto;
    list-style: none;
    padding: 0;
    margin: 0;
}

.feeds li {
    border-bottom: 1px dotted #ddd;
    width: auto;
}

.feeds li:last-child {
    border-bottom: none;
}

li {
    margin: 1.3rem 0;
    padding: 0 .5rem;
    position: relative;
}

li h1 {
    font-size: 1.2rem;
}

li h1 a {
    text-decoration: none;
}

li .buttons {
    float: right;
    position: relative;
    top: -2.3rem;
    opacity: 0;
    background: white;
}

li:hover .buttons {
    opacity: 1;
}
`
