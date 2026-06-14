import fstyle from "./quote.module.css"

export default function Quote({ quote }) {
    return (
        <div class="flex-container">
            <div className={fstyle.author}>
                {quote.author}
            </div>
            <div className={fstyle.text}>
                {quote.text}
            </div>
        </div>
    )
}