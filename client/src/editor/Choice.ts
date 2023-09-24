export default class Choice {

    static get toolbox() {
        return {
            title: 'Choice',
            icon: 'B'
        };
    }

    render() {
        const block = document.createElement('div');
        block.innerHTML = '<button contenteditable="true" class="my-1 inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2"></button>';
        return block;
    }

    save(){
        return {
        }
    }
}