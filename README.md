# Image Download Util

CLI tool built using Go to download images from a `.txt` file containing all the images addresses. 

## Installation

To install, follow these steps:

- Clone the repository: `git clone https://github.com/ivanatias/imgdl.git`
- Navigate to the cloned directory: `cd imgdl`
- Install the tool using Go: `go install`

### Important: 

You must install Go in your system. To install, follow the instructions on [Go's website](https://go.dev/doc/install).

## Usage 

Once installed, you can use the tool as follows: 

- Unix-like systems: `./imgdl -from /path/to/text.txt`
- Windows: `imgdl.exe -from C:\path\to\text.txt`

Optionally, you can provide a path to indicate where all the images should be saved. By default, all images will be saved on a folder called `imgdl` on the current working directory from where the CLI has been executed.

Providing a path where images should be saved:

`./imgdl -from /path/to/text.txt -to /path/where/images/are/saved`

Note that while the `to` flag is optional, `from` is required. 

You can run `./imgdl -help` to see all available options.

## Performance

![Speed test](./test.png)

It takes an average time of 30-35 seconds to download and save 600 images. However, this time can vary based on various factors such as:

- Network Speed: If the network speed is slow, it may take longer to download each image, resulting in increased overall download time.

- Image Size: Larger images will take more time to download and save compared to smaller ones. Therefore, if the images being downloaded have varying sizes, the download time will vary accordingly.

- Server Response Time: If the server from which the images are being downloaded has high response time, it can increase the overall download time.