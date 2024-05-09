
# tcpchat


**tcpchat** is a Go-based chat server that utilizes TCP connections to facilitate communication between users. It provides functionality for user registration, maintains a user IP map, allows the creation of groups, and efficiently routes messages to other clients.

![Untitled-2024-05-05-1539](https://github.com/ashupednekar/tcpchat/assets/25405037/c32ea4ba-028a-4a82-9395-64d31031843a)


## Features

- **User Registration**: Users can register themselves on the chat server.
- **User IP Map**: Maintains a mapping of users to their respective IP addresses.
- **Group Creation**: Users can create groups for specific discussions.
- **Message Routing**: Efficiently routes messages to appropriate clients based on their intended recipients.

## Design/Plan
Here's an overview of the overall design

> Following diagram illustrates the actions to be performed upon receiving messages conforming to these specific formats, and also covers the database design

<img width="1015" alt="Screenshot 2024-05-06 at 8 03 33 PM" src="https://github.com/ashupednekar/tcpchat/assets/25405037/bbcd4e52-4356-46ef-88c5-40147401529a">

> Here are the various channels to be used for communication across goroutines, and hence... users 😊

<img width="1015" alt="Screenshot 2024-05-09 at 12 55 40 AM" src="https://github.com/ashupednekar/tcpchat/assets/25405037/7b8a8ef9-5e85-46f5-8aca-458753740718">

## Demo

Here's a quick demo



https://github.com/ashupednekar/tcpchat/assets/25405037/8025a721-4556-40d7-bdfb-706bb6712c3f


https://www.veed.io/view/e3befc40-1007-44aa-80ba-bc3c7b364df8?panel=share

> Group Chat

<img width="1015" alt="Screenshot 2024-05-09 at 10 24 32 PM" src="https://github.com/ashupednekar/tcpchat/assets/25405037/890fee60-231d-495f-aeeb-ebb7b3c78836">


## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
