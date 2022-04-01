// TAG EDITOR
const tagEditor = new TagEditor({
    BlockSelectorName: "#b_TagEditor",
    TextBlockSelectorName: "#question__editor-tags",
});

// PREVIEW
// import "/static/js/showdown.js";
let mdConverter = new showdown.Converter();
let TabRadioButtons = document.getElementsByName("tab");

// Set Show and hide
TabRadioButtons.forEach((rb) => {
    const tabBlockName = rb.getAttribute("data-tab-block");
    const tabBlock = document.querySelector(tabBlockName);

    const tabBlocksName = rb.getAttribute("data-tab-blocks");
    const tabBlocks = document.querySelectorAll(tabBlocksName);

    let rb_OnChange = (e) => {
        tabBlocks.forEach((block) => {
            block.hidden = true;
        });

        if (rb.checked) {
            tabBlock.hidden = false;
        }
    }

    rb.addEventListener('change', rb_OnChange)
});

function CreateTag(name) {
    const btn_tag = document.createElement('div');
    btn_tag.setAttribute('class', 'btn-tag');
    btn_tag.innerHTML = name
    return btn_tag
}

// Вставка данных с Editor на Preview
const rb_tabPreview = document.getElementById("rb_tabPreview");
// EDITOR ELEMENTS
const editorTitle = document.getElementById("question__editor-title");
const editorTags = document.getElementById("question__editor-tags");
const editorText = document.getElementById("question__editor-text");
// PREVIEW ELEMETNS
const previewTitle = document.getElementById("question__preview-title");
const previewTags = document.getElementById("question__preview-tags");
const previewText = document.getElementById("question__preview-text");

let RenderPreveiw = () => {
    previewTitle.innerHTML = editorTitle.value
    // 
    previewTags.innerHTML = ""
    editorTags.value.split(" ").forEach((name) => {
        if (name != "") {
            previewTags.appendChild(CreateTag(name));
        }
    })
    // 
    previewText.innerHTML = mdConverter.makeHtml(editorText.value);
    console.log("changed")
}

rb_tabPreview.addEventListener("click", RenderPreveiw);


// var converter = new showdown.Converter()

//         let eText = document.getElementById("question__editor-text");
//         console.log(eText);
//         console.log(eText.value);
//         console.log(converter.makeHtml(eText.value));

//         let pText = document.getElementById("question__preview-text");
//         pText.innerHTML = converter.makeHtml(eText.value);
//         console.log(pText.innerHTML)