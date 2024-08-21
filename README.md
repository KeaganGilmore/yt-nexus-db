# yt-nexus-db

`yt-nexus-db` is a REST API built with Go and the Gin framework for managing and querying YouTube data. The API allows you to add and search for words, channels, and videos, as well as retrieve common words and top videos by channel.

## Features

- **Add and Manage Data**:
  - Add words to the dictionary.
  - Add YouTube channels.
  - Add video details and associated word counts.

- **Query Data**:
  - Retrieve the most common words used in a specific channel's videos.
  - Get the top videos from a channel.
  - Search for videos containing a specific keyword in a channel, across multiple channels, or in a list of video IDs.
  - Search for videos containing a specific keyword across the entire database.

## Installation

1. **Clone the Repository**:

    ```sh
    git clone https://github.com/yourusername/yt-nexus-db.git
    ```

2. **Navigate to the Project Directory**:

    ```sh
    cd yt-nexus-db
    ```

3. **Install Dependencies**:

    Ensure you have Go installed on your machine. Install the required Go modules:

    ```sh
    go mod tidy
    ```

4. **Build the Project**:

    ```sh
    go build -o yt-nexus-db
    ```

5. **Run the Project**:

    ```sh
    ./yt-nexus-db
    ```

    The server will start on port 8110.

## API Endpoints

### Add Word

- **POST** `/yt-nexus/dictionary`
- **Request Body**: `{ "word": "example" }`
- **Response**: `{ "word_id": 1 }`

### Add Channel

- **POST** `/yt-nexus/channel`
- **Request Body**: `{ "channel_name": "example_channel" }`
- **Response**: `{ "channel_id": 1 }`

### Add Video

- **POST** `/yt-nexus/video`
- **Request Body**: 
  ```json
  {
    "channel_id": 1,
    "video_id": "example_video",
    "word_counts": { "1": 10 }
  }
  ```
- **Response**: `{ "message": "Video added successfully!" }`

### Get Channel Common Words

- **GET** `/yt-nexus/channel/:channel_name/common-words`
- **Response**: `{ "common_words": { "example": 10 } }`

### Get Channel Top Videos

- **GET** `/yt-nexus/channel/:channel_name/top-videos`
- **Response**: `{ "top_videos": ["example_video_1", "example_video_2"] }`

### Get Videos With Keyword

- **GET** `/yt-nexus/channel/:channel_name/keyword/:keyword`
- **Response**: `{ "videos": [{ "video_id": "example_video", "count": 10 }] }`

### Search Across DB

- **GET** `/yt-nexus/search`
- **Query Parameter**: `keyword=example`
- **Response**: `{ "videos": [{ "video_id": "example_video", "count": 10 }] }`

### Search Across Channels

- **POST** `/yt-nexus/multi-channel-search`
- **Request Body**:
  ```json
  ["channel1", "channel2"]
  ```
- **Query Parameter**: `keyword=example`
- **Response**: `{ "videos": [{ "video_id": "example_video", "count": 10 }] }`

### Search Across Videos

- **POST** `/yt-nexus/multi-video-search`
- **Request Body**:
  ```json
  ["video1", "video2"]
  ```
- **Query Parameter**: `keyword=example`
- **Response**: `{ "videos": [{ "video_id": "example_video", "count": 10 }] }`

### About Us

Welcome to YT Nexus DB, a project deployed by LX Library (https://lxlibrary.online).

With the recent changes in YouTube's regulations, the previous YouTube Keyword Search (YTKS) setups became ineffective. In response, we’ve developed YT Nexus, this is a new database system designed to adapt to these new constraints. Accessed via the YT Nexus client website, this database serves as a comprehensive resource for video transcriptions and keyword searches.

We believe in the power of community contributions. You can help enhance YT Nexus DB by:

1. **Pulling the Repository**: Visit [YT-Nexus_API](https://github.com/KeaganGilmore/YT-Nexus_API) to get the latest scripts and tools for video transcription.
2. **Running the Script**: Use the provided script to transcribe YouTube videos.
3. **Contributing to the Database**: This script will automatically post to the deployed yt-nexus db allowing your contributions to be viewed by the entire community!

Due to restrictions on major traffic to YouTube for video transcriptions, it's crucial that multiple setups collaborate to make this project effective. Your contributions will help build a robust and comprehensive YouTube keyword search tool that benefits the entire community.

For more details, feel free to contact us at keagangilmore@gmail.com or reach out on Discord at keagan2980. Thank you for supporting YT Nexus DB!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

- **Email**: [keagangilmore@gmail.com](mailto:keagangilmore@gmail.com)
- **Discord**: [keagan2980](https://discord.com/users/keagan2980)
```
