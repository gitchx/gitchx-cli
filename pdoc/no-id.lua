function Pandoc(doc)
  for i, el in ipairs(doc.blocks) do
    if el.t == "Header" then
      el.identifier = ""
      if el.attributes then
        el.attributes["id"] = nil
      end
      doc.blocks[i] = el
    end
  end
  return doc
end

