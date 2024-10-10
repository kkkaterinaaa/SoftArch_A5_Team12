import React, { useState, useEffect } from 'react';
import MessageList from '../list/list.js';
import MessageForm from '../form/form.js';
import authService from '../../services/authService';

const HomePage = () => {
  const [messages, setMessages] = useState([]);

  const fetchLikesForMessage = async (messageId) => {
    const userId = authService.getUsername();
    try {
      const response = await fetch(`http://localhost:8070/likes/message/${messageId}?user_id=${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        }
      });
      const data = await response.json();
      // console.log("Fetched likes. Message: %s, Liked: %s, Username: %s", messageId, data.liked, username);
      return { likes: data.likes, likedBy: data.liked ? [userId] : [] };
    } catch (error) {
      console.error('Failed to fetch likes:', error);
      return { likes: 0, likedBy: [] };
    }
  };

  const fetchUsernameForUser = async (userId) => {
    try {
      const response = await fetch(`http://localhost:8070/users/${userId}`, {
        method: 'GET',
      });
      const data = await response.json();
      return data.Username;
    } catch (error) {
      console.error('Failed to fetch username:', error);
      return 'Unknown User';
    }
  };

  useEffect(() => {
    const fetchMessages = async () => {
      try {
        const response = await fetch('http://localhost:8070/messages', {
          method: 'GET',
        });
        const data = await response.json();

        const messagesWithLikes = await Promise.all(
          data.messages.map(async (message) => {
            const likesData = await fetchLikesForMessage(message.ID);
            const username = await fetchUsernameForUser(message.UserID);
            console.log(message.UserID);
            return { ...message, ...likesData, username };
          })
        );

        setMessages(messagesWithLikes.slice(0, 10));
      } catch (error) {
        console.error('Failed to fetch messages:', error);
      }
    };

    fetchMessages();
  }, []);

  const handleSendMessage = async (newMessage) => {
    const userId = authService.getUsername();

    if (!userId) {
      alert('You need to sign in to post a message.');
      return;
    }

    try {
      const response = await fetch('http://localhost:8070/messages', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: newMessage, user_id: userId }),
      });

      if (response.ok) {
        const result = await response.json();
        const likesData = await fetchLikesForMessage(result.message.id);
        const username = await fetchUsernameForUser(result.message.UserID);
        setMessages([{ ...result.message, ...likesData, username }, ...messages].slice(0, 10));
      } else {
        console.error('Failed to send message');
      }
    } catch (error) {
      console.error('Error while sending message:', error);
    }
  };

  const handleLikeMessage = async (id) => {
    const username = authService.getUsername();
    
    if (!username) {
      alert('You need to sign in to like a message.');
      return;
    }

    try {
      const response = await fetch('http://localhost:8070/likes', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ user_id: username, message_id: id }),
      });
  
      if (response.ok) {
        const updatedMessages = messages.map((message) => {
          if (message.ID === id) {
            const alreadyLiked = message.likedBy.includes(username);

            if (alreadyLiked) {
              return {
                ...message,
                likes: message.likes - 1,
                likedBy: message.likedBy.filter(user => user !== username),
              };
            } else {
              return {
                ...message,
                likes: message.likes + 1,
                likedBy: [...message.likedBy, username],
              };
            }
          }
          return message;
        });
        setMessages(updatedMessages);
      } else {
        console.error('Failed to like/unlike message');
      }
    } catch (error) {
      console.error('Error while liking/unliking message:', error);
    }
  };

  return (
    <div>
      <MessageList messages={messages} onLikeMessage={handleLikeMessage} />
      <MessageForm onSendMessage={handleSendMessage} />
    </div>
  );
};

export default HomePage;
