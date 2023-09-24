export default abstract class Block {

    protected wrapper;
    protected settingsWrapper;

    protected data: object;

    protected constructor() {
        this.wrapper = document.createElement('div');
        this.settingsWrapper = document.createElement('div');
    }

    abstract make();

    abstract makeSettings();

    render() {
        this.save(this.wrapper);
        this.make();
        return this.wrapper;
    }

    renderSettings() {
        const self = this;
        this.settingsWrapper.innerHTML = '';
        self.makeSettings().forEach(tune => {
            let button = document.createElement('div');
            button.innerHTML = '<div class="ce-popover-item"><div class="ce-popover-item__icon">' + tune.icon + '</div><div class="ce-popover-item__title">' + tune.label + '</div></div>'
            button.addEventListener('click', function () {
                tune.onActivate();
                self.render();
                // if (self.refresh != null) self.refresh();
            })
            self.settingsWrapper.appendChild(button);
        });
        return this.settingsWrapper;
    }

    abstract save();

    abstract refresh();

    abstract static get toolbox();
}