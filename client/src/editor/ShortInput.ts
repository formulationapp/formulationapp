export default class ShortInput {

    static get toolbox() {
        return {
            title: 'Short input',
            icon: 'aa'
        };
    }

    render() {
        const block = document.createElement('div');
        block.innerHTML = '<input type="text" class="flex my-4 h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50">';
        return block;
    }

    save() {
        return {}
    }
}