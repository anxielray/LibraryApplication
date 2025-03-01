# Anxiel Archives

Anxiel Archives is a web-based learning companion platform that helps users track their reading progress, manage their personal library, and participate in book clubs. The platform is designed to enhance the reading and learning experience through organized tracking and community engagement.

## Features

- **Personal Library Management**: Track and manage your reading collection
- **Reading Progress Tracking**: Monitor your reading habits and progress
- **Book Clubs**: Join and participate in community discussions
- **User Authentication**: Secure login and registration system
- **Responsive Design**: Mobile-friendly interface using Bootstrap

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS, JavaScript
- **CSS Framework**: Bootstrap 5
- **Icons**: Font Awesome
- **Fonts**: Playfair Display, Roboto

## Getting Started

### Prerequisites

- Go 1.x or higher
- Web browser
- Internet connection (for CDN resources)

### Installation

1. Clone the repository
```bash
git clone https://github.com/anxielray/LibraryApplication.git
```

2. Navigate to the project directory
```bash
cd LibraryApplication
```

3. Run the server
```bash
go run main.go
```

4. Access the application at `http://localhost:21112`

## Project Structure

```
LibraryApplication/
├── auth/
│   └── authentication handlers
├── library/
│   └── library management
├── profile/
│   └── user profile handlers
├── templates/
│   ├── index.html
│   ├── login.html
│   ├── registration.html
│   └── dashboard.html
└── main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Anxiel Archives - [@Anxiel](https://github.com/anxielray)

Project Link: [https://github.com/anxielray/anxiel-archives](https://github.com/anxielray/LibraryApplication)