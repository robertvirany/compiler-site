document.addEventListener('DOMContentLoaded', () => {
    const compilerSelect = document.getElementById('compiler');
    const modeSelect = document.getElementById('mode');

    const modeOptions = {
        gcc: [
            { value: 'preprocess', label: 'Preprocess' },
            { value: 'asm', label: 'Assembly' },
            { value: 'bin', label: 'Executable' },
        ],
        clang: [
            { value: 'preprocess', label: 'Preprocess' },
            { value: 'ir', label: 'LLVM IR' },
            { value: 'asm', label: 'Assembly' },
            { value: 'bin', label: 'Executable' },
        ]
    }

    function updateModeOptions() {
        const compiler = compilerSelect.value;
        modeSelect.innerHTML = ''; // wipe existing

        for (const opt of modeOptions[compiler]) {
            const o = document.createElement('option');
            o.value = opt.value;
            o.textContent = opt.label;
            modeSelect.appendChild(o);
        }
    }

    compilerSelect.addEventListener('change', updateModeOptions);
    updateModeOptions(); //initialize on page load
});