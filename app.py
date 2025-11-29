from flask import Flask, render_template, jsonify

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/api/ping')
def ping():
    return jsonify({"message": "pong"})

if(__name__ == "__main__"):
    app.run(debug=True, host='0.0.0.0')
# ...existing code...
