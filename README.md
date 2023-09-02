![Logo-white](https://user-images.githubusercontent.com/11553349/144132818-8bbcd9ba-fa0d-4a99-aa80-9fea8afcefe3.png)

[Licensephobia.com](https://licensephobia.com) is a tool that let's you easily search for NPM (and soon pip) packages to view the licenses and get a license summary about your apps.

It offers a public user interface, a REST API as well as a CLI. All of the libraries/tools are open source and built for the open source community!

## How it's done
- We maintain a Postgres Database (hosted on Supabase) from which we are querying data about what licenses there are, what usage conditions they have and where to find further information about them. The data is continuously fetched from [SPDX data](https://spdx.org/licenses/) as well as a lot of manual screening of licenses.
- We use the public NPM registry API to search for packages and the license info, but the architecture is set up in a way that let's us add other package managers in the future.
- Users can view license info on specific packages on our UI, or upload/scan package.json files to get a summary of all the licenses used.

As said, we will open-source all the packages/apps we build, at the the stuff we build will make up our service architecture:
<p align="center">
<img src="https://user-images.githubusercontent.com/11553349/144129612-38f126d7-8ac8-4be5-81d8-326e89f265c6.png" alt="Service architecture" style="height: 400px;">
</p>

## Long-tearm goals:
- Support many different package managers. The list, right now, includes:
  - Pip (requirements.txt, Pipfile, Pyproject.tom etc)
  - Cargo/Rust
  - Go
  - NuGet
  - ...
- A robust CLI that integrates with CI/CD tools to continuously scan packages and provides warnings/infrmation
- The web's biggest list of machine-readable license condition data

## Why not use [choosealicense.com](https://choosealicense.com/)?
Well, right now you should probably use this if you want to know more about open source licenses! But there are some key differences to licensephobia:
- It only supports a selected number of licenses
- You can only get information about licenses, not packages or your package dependency files.
- It does not offer a CLI (as far as I know?)

## FAQ
- **F: Does Licensephobia provide any legal guarantee about the license summaries?**
  - **A: No!** While we try our best to provide an accurate and correct representation of the licenses and take care to only show valid information in the summary, there might be mistakes on our side or in the data we use, that lead to inconsistancies. Because of this, we DO NOT provide any legal advice and don't guarantee the correctness of our data. If you find any mistake on our side, please open an issue on Github!
- **F: Does Licensephobia use the data uploaded via package files in any other way, apart from the license stuff?**
  - **A: Of course not!** Uploaded files are not saved and only the package names and versions are read by our sourcecode. In fact, you can just browse the sourcecode of this repo to verify this. However, the whole uploaded file will be sent to our server, just to make the parsing of files easier and faster.
- **F: Is the name `Licensephobia` inspired by some other tool?**
  - **A: Oh yes,** it's basically a rip-off of the name of the amazing tool [Bundlephobia](https://bundlephobia.com). While we don't have any relation to the Bundlephobia maintainers, we recommend to check them out!
