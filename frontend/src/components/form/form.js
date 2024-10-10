import React, { useState } from 'react';
import './form.css';  // Import the new CSS file

const MessageForm = ({ onSendMessage }) => {
  const [message, setMessage] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();

    if (message.trim() === '') {
      return; 
    }

    onSendMessage(message);
    setMessage(''); 
  };

  return (
    <form onSubmit={handleSubmit} className="message-form">
      <textarea
        className="message-input"
        placeholder="Enter your message..."
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        maxLength={400} 
      />
      <button type="submit" className="submit-button">Send Message</button>
    </form>
  );
};

export default MessageForm;
