import Block from "@/editor/Block";

export default class ShortAnswer extends Block {

    constructor({data}) {
        super();
        if (Object.keys(data).length > 0) this.data = data;
    }

    protected data = {
        text: '',
        label: '',
        required: false,
    };

    make() {
        const block = document.createElement('div');
        block.innerHTML = '<span contenteditable="true" class="font-semibold" placeholder="Label"></span><span></span>';
        if (this.data.required)
            block.innerHTML += '<div class="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent bg-primary text-primary-foreground hover:bg-primary/80 ml-2 ml-2">*</div>';

        block.innerHTML += '<input type="text" class="t flex my-4 h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50">';
        if (this.data.text != null)
            block.querySelector('input').setAttribute('value', this.data.text);
        if (this.data.label != null)
            block.querySelector('span').innerText = this.data.label;
        this.wrapper.innerHTML = block.innerHTML;
    }

    makeSettings() {
        const self = this;
        return [
            {
                icon: `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
      <path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M6 7L6 12M6 17L6 12M6 12L12 12M12 7V12M12 17L12 12"></path>
      <path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M16 11C16 10.5 16.8323 10 17.6 10C18.3677 10 19.5 10.311 19.5 11.5C19.5 12.5315 18.7474 12.9022 18.548 12.9823C18.5378 12.9864 18.5395 13.0047 18.5503 13.0063C18.8115 13.0456 20 13.3065 20 14.8C20 16 19.5 17 17.8 17C17.8 17 16 17 16 16.3"></path>
    </svg>`,
                label: 'Required',
                onActivate() {
                    self.data.required = !self.data.required;
                }
            }
        ];
    }

    static toolbox() {
        return {
            title: 'Short input',
            icon: 'a'
        };
    }

    save(blockContent) {
        try {
            this.data.label = this.wrapper.querySelector('span').innerText;
            this.data.text = this.wrapper.querySelector('input').value;
            return this.data;
        } catch (err) {
        }
    }
}