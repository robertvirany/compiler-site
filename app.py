#%%
from flask import Flask, request, jsonify, send_from_directory
import tempfile
import subprocess
import os
import binascii

#%%
app = Flask(__name__, static_folder='public', static_url_path='')

@app.route('/')
def index():
    return send_from_directory(app.static_folder, 'index.html')

@app.route('/compile', methods=['POST'])
def compile_code():
    data = request.get_json()
    code = data.get('code')
    compiler = data.get('compiler')
    mode = data.get('mode')

    with tempfile.NamedTemporaryFile(delete=False, suffix='.c') as tmp:
        tmp.write(code.encode())
        tmp_path = tmp.name

    output = ''
    err = ''
    cmd = []

    try:
        if compiler == 'gcc':
            if mode == 'preprocess':
                cmd = ['gcc', '-E', tmp_path, '-o', '-']
            elif mode == 'asm':
                cmd = ['gcc', '-S', tmp_path, '-o', '-']
            elif mode == 'bin':
                obj_path = tmp_path.replace('.c', '.o')
                cmd = ['gcc', '-c', tmp_path, '-o', obj_path]
            else:
                return jsonify({'error': 'unsupported mode', 'output': ''}), 400

        elif compiler == 'clang':
            if mode == 'preprocess':
                cmd = ['clang', '-E', tmp_path, '-o', '-']
            elif mode == 'ir':
                cmd = ['clang', '-S', '-emit-llvm', tmp_path, '-o', '-']
            elif mode == 'asm':
                cmd = ['clang', '-S', tmp_path, '-o', '-']
            elif mode == 'bin':
                cmd = ['clang', '-c', tmp_path, '-o', '-']
            else:
                return jsonify({'error': 'unsupported mode', 'output': ''}), 400
        if mode == 'bin' and compiler == 'clang':
            output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
            output = binascii.hexlify(output).decode()
        elif mode == 'bin' and compiler == 'gcc':
            output = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
            with open(obj_path, 'rb') as f:
                output = binascii.hexlify(f.read()).decode()
        else:
            output = subprocess.check_output(cmd, stderr=subprocess.STDOUT).decode()
    except subprocess.CalledProcessError as e:
        output = e.output.decode()
        err = str(e)
    finally:
        os.remove(tmp_path)
        if mode == 'bin' and 'obj_path' in locals():
            os.remove(obj_path)

    return jsonify({'output': output, 'error': err})

#%%
if __name__ == '__main__':
    app.run(host='127.0.0.1', port=8080)
