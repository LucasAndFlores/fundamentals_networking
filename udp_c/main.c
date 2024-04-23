#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
    int port = 5501;
    int sockfd;
    struct sockaddr_in myaddr, remoteAddr;

    char buffer[1024];
    socklen_t addr_size;

    sockfd = socket(AF_INET, SOCK_DGRAM, 0);
    
    memset(&myaddr, '\0', sizeof(myaddr));
    myaddr.sin_family = AF_INET;
    myaddr.sin_port = htons(port);
    myaddr.sin_addr.s_addr = inet_addr("127.0.0.1");
    
    bind(sockfd, (struct sockaddr*)&myaddr, sizeof(myaddr));
    while (1) { // Loop to continuously receive messages
        addr_size = sizeof(remoteAddr);
        ssize_t recv_len = recvfrom(sockfd, buffer, 1024, 0, (struct sockaddr*)&remoteAddr, &addr_size);
        if (recv_len > 0) {
            buffer[recv_len] = '\0'; // Null-terminate the received data
            printf("Received data from %s:%d: %s\n", inet_ntoa(remoteAddr.sin_addr), ntohs(remoteAddr.sin_port), buffer);
        } else {
            fprintf(stderr, "Failed to receive data\n");
        }
    }

    close(sockfd);

    return 0;
}
