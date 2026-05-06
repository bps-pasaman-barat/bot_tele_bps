import re
import argparse
from pathlib import Path

from langchain_community.document_loaders import PyPDFLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter


def clean_text(text):
    """
    Membersihkan header/footer yang berulang agar RAG lebih akurat.
    """
    text = re.sub(r'https://pasamanbaratkab\.bps\.go\.id', '', text)
    text = re.sub(r'Keadaan Angkatan Kerja Kabupaten Pasaman Barat 2024', '', text)
    text = re.sub(r'\s+', ' ', text).strip()
    return text


def parse_pdf_for_rag(file_path):
    print(f"1. Memuat PDF: {file_path}...")
    
    loader = PyPDFLoader(file_path)
    documents = loader.load()

    print(f"   Berhasil memuat {len(documents)} halaman.\n")

    print("2. Membersihkan teks dari header/footer...")
    for doc in documents:
        doc.page_content = clean_text(doc.page_content)

        doc.metadata['source'] = 'BPS Pasaman Barat'
        doc.metadata['year'] = '2024'

    print("3. Memotong teks (Chunking) untuk RAG...")

    text_splitter = RecursiveCharacterTextSplitter(
        chunk_size=1000,
        chunk_overlap=150,
        length_function=len,
        separators=["\n\n", "\n", ".", " ", ""]
    )

    chunks = text_splitter.split_documents(documents)

    print(f"   Berhasil memotong PDF menjadi {len(chunks)} chunks.\n")

    return chunks


def save_markdown(chunks, output_file):
    print(f"4. Menyimpan hasil ke '{output_file}'...")

    with open(output_file, "w", encoding="utf-8") as f:

        f.write("# HASIL PARSING & CHUNKING DOKUMEN\n\n")
        f.write(f"**Total Chunks:** {len(chunks)}\n\n")
        f.write("---\n\n")

        for i, chunk in enumerate(chunks):

            halaman = chunk.metadata.get('page', 'Tidak diketahui')

            f.write(f"## CHUNK #{i+1} (Halaman {halaman})\n\n")
            f.write(f"**Metadata:** `{chunk.metadata}`\n\n")
            f.write(f"{chunk.page_content}\n\n")
            f.write("---\n\n")

    print("   Selesai!")


def main():

    parser = argparse.ArgumentParser(
        description="Parse PDF menjadi chunk Markdown untuk RAG"
    )

    parser.add_argument(
        "pdf_path",
        help="Path file PDF yang akan diparse"
    )

    parser.add_argument(
        "output",
        nargs="?",
        help="Nama file output markdown (opsional)"
    )

    args = parser.parse_args()

    pdf_path = args.pdf_path

    if args.output:
        output_file = args.output
    else:
        # otomatis buat nama output dari nama pdf
        output_file = Path(pdf_path).stem + ".md"

    chunks = parse_pdf_for_rag(pdf_path)

    save_markdown(chunks, output_file)


if __name__ == "__main__":
    main()