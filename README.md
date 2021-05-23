# AboutPage
A web application for viewing information about the contents of a web page

# Here's how it works
Let's imagine you're developing or using a web application, and you want to know important information about the contents of a web page such as:
* HTML Version
* Page Title
* Headings count by level
* Amount of internal and external links
* Amount of inaccessible links
* If a page contains a login form

You can enter the URL of the web page you want information about, and AboutPage will display a dashboard of this information.

# Use cases and why it is useful

A large number of DOM elements can impact web page performance when they are manipulated using JavaScript. Searching for an ID in a page with 40,000 elements is going to be slower than a page with only 400. Changing a CSS style which is used by most of the page will likely lead to more unecessary layout than it would on a simpler page [1].

According to a [Google web page performance audit][2], large DOM sizes should be avoided:

> A large DOM tree can slow down your page performance in multiple ways:
>
> * Network efficiency and load performance
>
>   A large DOM tree often includes many nodes that aren't visible when the user first loads the page, which unnecessarily increases data costs for your users and slows down load time.
>
>* Runtime performance
>
>   As users and scripts interact with your page, the browser must constantly recompute the position and styling of nodes. A large DOM tree in combination with complicated style rules can severely slow down rendering.
>
> * Memory performance
>
>   If your JavaScript uses general query selectors such as document.querySelectorAll('li'), you may be unknowingly storing references to a very large number of nodes, which can overwhelm the memory capabilities of your users' devices.

https://web.dev/dom-size/#how-to-optimize-the-dom-size

[Google Lighthouse][2] says that an optimal DOM tree:

- Has less than 1,500 nodes total.
- Has a depth less than 32 nodes.
- Has a parent node with less than 60 child nodes.

As a developer, the audit provided by AboutPage can help you better design your websites to improve their quality, and as a user this information can help you get an idea which open web pages are consuming more of your ram and slowing down your device. It can also help developers keep track of how big a web page's DOM tree is getting as well as monitor the size of the DOM tree during DOM tree size optimization.

You can even use this information in web content analytics to compare webpages of different websites and find trends among them, answering questions such as do ecommerce websites make use of external links more than question and answer websites? The audit can also reveal insights about a webpage such as a large number of inaccessible links indicating heavy API usage or query string parameters in URLs.

# Instructions to run the application
Open a terminal and run the following commands:
```
//clone repository to your local computer
$ git clone https://github.com/ccaneke/aboutpage

// cd into the aboutpage directory
$ cd aboutpage

// compile the main package to generate an executable
$ go build aboutpage.go

// run the executable (on Windows run aboutpage without the "./" instead)
$ ./aboutpage
```

[1]: https://kellegous.com/j/2013/01/26/layout-performance/
[2]: https://web.dev/dom-size/#how-the-lighthouse-dom-size-audit-fails
 