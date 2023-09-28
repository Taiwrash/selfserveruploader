## SelfServerUploader - Easy File Uploads with Fiber

SelfServerUploader is a Go package designed to simplify the process of handling file uploads in your web applications. It's built on top of the Fiber web framework, making it easy to receive, save, and provide access to uploaded files.

## Table of Contents
- [SelfServerUploader - Easy File Uploads with Fiber](#selfserveruploader---easy-file-uploads-with-fiber)
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Just Before You Run](#just-before-you-run)
- [In your HTML file](#in-your-html-file)
- [To use as API](#to-use-as-api)
- [Good to go!](#good-to-go)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Handling file uploads is a common task in web development, whether you're creating a platform for sharing images, managing documents, or any other application that requires users to upload files. SelfServerUploader streamlines this process and helps you get started quickly.

## Installation

To use SelfServerUploader in your Go application, follow these steps to install it:

```bash
go get github.com/Taiwrash/selfserveruploader
```

## Usage

Here's how you can use the SelfServerUploader package in your Fiber application:

1. Import the package:

   ```go
   import (
       "github.com/Taiwrash/selfserveruploader"
       "github.com/gofiber/fiber/v2"
   )
   ```

2. Set up a Fiber app:

   ```go
   app := fiber.New()
   ```

3. Use the `selfserveruploader.Upload` function as a middleware for handling file uploads:

   ```go
   app.Post("/upload", selfserveruploader.Upload)
   ```

   In this example, the `/upload` route is configured to handle file uploads.

4. Customize the upload behavior by modifying the `Upload` function as needed. You can specify file naming conventions, save locations, and more.

## Example

Here's a complete example of how to use the SelfServerUploader package in your Fiber application:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Taiwrash/selfserveruploader"
)

func main() {
    app := fiber.New()

    // Handle file uploads at the "/upload" endpoint
    app.Post("/upload", selfserveruploader.Upload)

    app.Listen(":4000")
}
```
## Just Before You Run
- Create a folder named `uploads` in the same directory as your `main.go` file.
- This is where the uploaded files will be saved.

## In your HTML file
- Create a form with the `enctype` attribute set to `multipart/form-data`.
- Create an input field with the `type` attribute set to `file`.
- Set the `name` attribute of the input field to `file`.
- Set the `action` attribute of the form to the route you specified in your Fiber application.
- Set the `method` attribute of the form to `POST`.
- Set the `enctype` attribute of the form to `multipart/form-data`.

```html
<form action="/upload" method="POST" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="submit" value="Upload">
</form>
```

## To use as API
- Select form data as `form-data` in Postman.
- Set the `key` to `file` with title `image`.
- Set the `value` to the file you want to upload.

## Good to go!
Run the above code, and your Fiber application will be able to receive and save uploaded files.

## Contributing

We encourage contributions from the community! If you have ideas for improvements, bug reports, or want to contribute code, please visit the GitHub repository at [https://github.com/Taiwrash/selfserveruploader](https://github.com/Taiwrash/selfserveruploader). We welcome your feedback and contributions.

## License

This package is distributed under the MIT License. See the [LICENSE](LICENSE) file for details.

Feel free to customize this README to include additional information or usage examples specific to your package. Happy uploading! 🚀