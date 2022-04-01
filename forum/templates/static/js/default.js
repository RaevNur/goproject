// https://github.com/Dias1c/frontend-utils-js
// TagEditor - Tag Editor Util
class TagEditor {
    // ! HTML Elements
    B_TagEditor; // new HTMLDivElement();
    Tb_Input; // new HTMLInputElement();
    Tb_Tags; // new HTMLInputElement();

    // ! SETTIGNS Default
    SBlockSelectorName = '#b_TagEditor';            // visible TagEditor block SelectorName
    STextBlockSelectorName = '#tb_TagEditor';       // input SelectorName
    SHasDoubles = false;                            // AddDoubles?
    SToLower = true;                                // ToLowercase Tag?
    SMaxTags = 0;                                   // Max Tags Count, 0 == Unlimited
    STags = [];                                     // Writed Tags
    SSeparator = ' ';                               // Split by separator

    // Inits TagEditor By Settings
    constructor(settings) {
        // Set Settings
        this.init_Settings(settings);
        // Init Inputs and Blocks on html
        this.B_TagEditor = document.querySelector(this.SBlockSelectorName);
        this.Tb_Tags = document.querySelector(this.STextBlockSelectorName);
        this.init_Tb_Input();
        // Set Tags On input and STags
        this.RefreshTags();
    }
    // Set Default Settings by Settings
    init_Settings(settings) {
        this.SBlockSelectorName = (settings.BlockSelectorName) ? settings.BlockSelectorName : this.SBlockSelectorName;
        this.STextBlockSelectorName = (settings.TextBlockSelectorName) ? settings.TextBlockSelectorName : this.STextBlockSelectorName;
        this.SHasDoubles = (settings.HasDoubles) ? settings.HasDoubles : this.SHasDoubles;
        this.SToLower = (settings.ToLower) ? settings.ToLower : this.SToLower;
        this.SMaxTags = (settings.MaxTags) ? settings.MaxTags : this.SMaxTags;
        this.STags = (settings.Tags) ? settings.Tags : this.STags;
    }
    init_Tb_Input() {
        if (this.B_TagEditor.querySelector('input') == null) {
            this.B_TagEditor.appendChild(document.createElement('input'));
        }
        this.Tb_Input = this.B_TagEditor.querySelector('input');

        // Add Events
        this.Tb_Input.addEventListener('keyup', (e) => {
            if (e.key === ' ') {
                this.Tb_Input.value = '';
            }
        });
        this.Tb_Input.addEventListener('keydown', (e) => {
            if (e.key === ' ' || e.key === 'Enter') {
                let tagname = this.Tb_Input.value;
                this.AddTag(tagname);
                this.Tb_Input.value = '';
            } else if (e.key === 'Backspace' && this.Tb_Input.value == "") {
                let tagname = this.RemoveLastTag();
                this.Tb_Input.value = tagname + " ";
            }
        });

        // Add Tags
        this.AddTag(this.Tb_Input.value);
        this.Tb_Input.value = '';
    }

    // ! Public Methods
    // Check Tagname for valid
    IsValidName(name) {
        if (this.SToLower) {
            name = name.toLowerCase();
        }
        if (name == "") {
            return false
        } else if (this.SMaxTags != 0 && this.STags.length >= this.SMaxTags) {
            return false
        } else if (!this.SHasDoubles && this.STags.includes(name)) {
            return false
        }
        return true
    }
    // Add Tag to TagEditor
    AddTag(name) {
        name.split(this.SSeparator).forEach((tagname, index) => {
            if (!this.IsValidName(tagname)) {
                return
            }
            if (this.SToLower) {
                tagname = tagname.toLowerCase();
            }
            this.STags.push(tagname);
        });
        this.RefreshTags();
    }
    // Removes last tag from TagEditor
    RemoveLastTag() {
        let name = this.STags.pop();
        this.RefreshTags();
        return (name != null) ? name : "";
    }
    // Removes tag by name from TagEditor
    RemoveTag(name) {
        this.STags = this.STags.filter(function (value, index, arr) {
            return value != name;
        });
        this.RefreshTags();
        return (name != null) ? name : "";
    }
    // Sets Tags to TagEditor by STags
    RefreshTags() {
        while (this.B_TagEditor.firstChild != this.Tb_Input) {
            this.B_TagEditor.removeChild(this.B_TagEditor.firstChild);
        }
        this.STags.slice().reverse().forEach((tagname, index) => {
            const newtag = TagEditor.CreateTag(tagname, () => {
                this.RemoveTag(tagname);
            });
            this.B_TagEditor.prepend(newtag);
        });
        this.Tb_Tags.value = this.STags.join(' ');
    }

    // Clears all Tags
    Clear() {
        this.STags = [];
        this.RefreshTags();
    }

    // ! Static Methods

    // CreateTag - returns (tag) HTMLDivElement
    // name - tag name
    // btn_remove_click - function wich work on click remove tag
    static CreateTag(name, btn_remove_click) {
        const btn_tag = document.createElement('div');
        btn_tag.setAttribute('class', 'btn-tag');
        const tag_name = document.createElement('span');
        tag_name.innerHTML = name
        const btn_remove = document.createElement('span');
        btn_remove.setAttribute('class', 'remove');
        btn_remove.innerHTML = '×';
        btn_remove.addEventListener(
            'click',
            (btn_remove_click != null) ? function () { btn_remove_click(); btn_tag.remove(); } : () => { btn_tag.remove(); }
        );
        // Construct tag
        btn_tag.appendChild(tag_name);
        btn_tag.appendChild(btn_remove);
        return btn_tag
    }
}