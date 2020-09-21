#URL Shortener API
This a simple URL shortener API which uses custom base62 conversion to generate unique short links.
####Tech Stacks
- Golang
- MongoDB

####Usage
After downloading the project you should specify your MongoDB's URL and some other constants inside `.env` file.

### API Documentation
> API endpoint: http://localhost:3000/api/v1/links
#### API Payload

- "URL" - Original URL(String, Mandatory)
- "userid" - UserID(Integer, Optional)

#### API Response
    status 200
    {
        "url": Original URL,
        "short_url": Custom short URL
    }
    
    status 400(Invalid json body's format)
    {
        "error": "cannot parse JSON",
        "title": "Invalid request format",
        "type": "Invalid request"
    }
    
    status 400(Invalid payload)
    {
      "error": "URL: %Payload URL% does not validate as url",
      "title": "There was a validation error",
      "type": "Validation error"
    }
    
    