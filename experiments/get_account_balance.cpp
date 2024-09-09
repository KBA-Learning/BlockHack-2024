#include <iostream>
#include <string>
#include <curl/curl.h>

// Callback function to handle data received from the server
static size_t write_callback(void *contents, size_t size, size_t nmemb, std::string *buffer)
{
    size_t total_size = size * nmemb;
    buffer->append((char *)contents, total_size); // Append data to the buffer
    return total_size;
}

int main()
{
    CURL *curl;
    CURLcode res;
    std::string response_data;

    // JSON data to send in POST request
    std::string json_data = R"({"jsonrpc":"2.0","method":"eth_getBalance","params":["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "latest"],"id":1})";

    // Initialize libcurl
    curl_global_init(CURL_GLOBAL_DEFAULT);
    curl = curl_easy_init();

    if (curl)
    {
        // Set the URL for the request
        curl_easy_setopt(curl, CURLOPT_URL, "http://localhost:8545");

        // Set the request to POST
        curl_easy_setopt(curl, CURLOPT_POST, 1L);

        // Set the JSON data as the POST fields
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, json_data.c_str());

        // Set the Content-Type to application/json
        struct curl_slist *headers = NULL;
        headers = curl_slist_append(headers, "Content-Type: application/json");
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

        // Set the callback function to handle response data
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);

        // Set a pointer to the response string to be filled by the callback
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_data);

        // Perform the HTTP request
        res = curl_easy_perform(curl);

        // Check for errors
        if (res != CURLE_OK)
        {
            std::cerr << "curl_easy_perform() failed: " << curl_easy_strerror(res) << std::endl;
        }
        else
        {
            // Print the response data
            std::cout << "Response data: " << response_data << std::endl;
        }

        // Cleanup
        curl_slist_free_all(headers); // Free the list of headers
        curl_easy_cleanup(curl);
    }

    // Cleanup global libcurl resources
    curl_global_cleanup();

    return 0;
}
