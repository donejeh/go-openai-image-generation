# OpenAI Image Generator

This project generates images based on user prompts using the OpenAI API.

## Prerequisites

- Go 1.23.3 or later
- An OpenAI API key

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/openai-images.git
    cd openai-images
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Create a [.env](http://_vscodecontentref_/0) file in the root directory of the project with your OpenAI API key:

    ```sh
    echo "OPENAI_API=your_openai_api_key" > .env
    ```

    Replace `your_openai_api_key` with your actual OpenAI API key.

## Running the Project

1. Run the Go program:

    ```sh
    go run main.go
    ```

2. Enter your prompt when prompted:

    ```sh
    Enter your prompt: >
    ```

3. The generated image will be saved as `image.png` in the root directory.

## License

This project is licensed under the MIT License.