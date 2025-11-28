import flask 

app = flask.Flask(__name__)

@app.route('/login')
def login():
    return "Hello World"

if(__name__ == "__main__"):
    app.run(debug=True)

#Hello world
#I LOVE NAOMI
