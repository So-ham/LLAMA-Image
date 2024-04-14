import ollama
from flask import Flask, request, jsonify

app = Flask(__name__)

# Define route for handling chat messages
@app.route('/chat', methods=['POST'])
def chat_handler():
    # Get user message from request
    request_data = request.json
    user_messages = request_data.get('messages', [])
    print(user_messages)

    # Send user messages to Ollama chat
    response = ollama.chat(model='llama2', messages=user_messages)
    message_content = response['message']['content']
    
    return jsonify({'response': message_content})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
