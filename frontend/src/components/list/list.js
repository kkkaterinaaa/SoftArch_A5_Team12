import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHeart } from '@fortawesome/free-solid-svg-icons';
import authService from '../../services/authService.js';
import './list.css';

const MessageList = ({ messages, onLikeMessage }) => {
  const username = authService.getUsername();

  return (
    <div className="message-box">
      {/* Scrollable messages list */}
      <div className="message-list">
        {messages.map((message) => (
          <div key={message.ID} className="message-container">
            <p className="message-username">{message.username}</p>
            <p className="message-text">{message.Content}</p>

            <button className="like-button" onClick={() => onLikeMessage(message.ID)}>
              <span className="heart" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                <FontAwesomeIcon icon={faHeart} style={{ color: message.likedBy.includes(username) ? 'red' : 'gray' }} />
                <span>{message.likes}</span>  
              </span>
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default MessageList;
